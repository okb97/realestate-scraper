# ローカル実行安定性向上戦略

## 🚨 ローカル実行の課題

### 主要リスク
1. **PC電源OFF**: 手動シャットダウン、電力不足
2. **スリープモード**: macOSの自動スリープ
3. **アプリ終了**: 誤操作やシステム更新
4. **ネットワーク断**: WiFi切断、ISP障害
5. **システム更新**: macOS自動更新による再起動

## 💡 安定性向上ソリューション

### 🔋 Option 1: MacBook安定化設定（最も現実的）

#### システム設定最適化
```bash
# 1. 電源管理設定
sudo pmset -a sleep 0                    # スリープ無効
sudo pmset -a displaysleep 0             # ディスプレイスリープ無効  
sudo pmset -a disksleep 0               # ディスクスリープ無効
sudo pmset -a powernap 0                # Power Nap無効
sudo pmset -a standby 0                 # スタンバイ無効

# 2. 自動更新無効
sudo softwareupdate --schedule off      # 自動ソフトウェア更新OFF

# 3. Caffeinate（システム起動維持）
caffeinate -s &                         # スリープ防止デーモン起動
```

#### 自動復旧機能付きスクリプト
```bash
#!/bin/bash
# robust_daily_scraping.sh

LOG_DIR="./logs/$(date +%Y%m%d)"
LOCK_FILE="/tmp/scraper.lock"
PID_FILE="/tmp/scraper.pid"

# ロックファイルチェック（重複実行防止）
if [ -f "$LOCK_FILE" ]; then
    echo "既に実行中です。終了します。"
    exit 1
fi

# ロックファイル作成
echo $$ > "$LOCK_FILE"
echo $$ > "$PID_FILE"

# 終了時クリーンアップ
trap 'rm -f "$LOCK_FILE" "$PID_FILE"; exit' INT TERM EXIT

# システム状態チェック
check_system() {
    # ディスク容量チェック
    AVAILABLE=$(df -h . | awk 'NR==2 {print $4}' | sed 's/Gi//')
    if [ "$AVAILABLE" -lt 50 ]; then
        echo "警告: ディスク容量不足 (${AVAILABLE}GB残り)"
        # Slackやメール通知を送信
        send_notification "ディスク容量不足: ${AVAILABLE}GB"
    fi
    
    # メモリ使用量チェック
    MEMORY_USAGE=$(vm_stat | grep "Pages free" | awk '{print $3}' | sed 's/.$//')
    if [ "$MEMORY_USAGE" -lt 100000 ]; then
        echo "警告: メモリ不足"
    fi
}

# 実行状況を記録
update_status() {
    echo "$(date): $1" >> "./logs/execution_status.log"
    # 外部監視システムに状況送信（オプション）
    curl -X POST "https://your-monitoring-service.com/status" \
         -d "status=$1&timestamp=$(date +%s)" 2>/dev/null || true
}

# メイン処理
main() {
    update_status "スクレイピング開始"
    
    for region_id in {1..10}; do
        update_status "地域${region_id}処理開始"
        
        # 3回までリトライ
        for attempt in {1..3}; do
            if timeout 7200 go run cmd/scraper/main.go -area=$region_id; then
                update_status "地域${region_id}処理完了"
                break
            else
                echo "地域${region_id} 試行${attempt}回目失敗"
                if [ $attempt -eq 3 ]; then
                    update_status "地域${region_id}処理失敗（3回試行後）"
                    send_notification "地域${region_id}の処理が失敗しました"
                else
                    sleep 1800  # 30分待機後リトライ
                fi
            fi
        done
        
        # システム状態チェック
        check_system
        
        # 地域間休憩
        sleep 30
    done
    
    update_status "全地域処理完了"
}

# 通知機能
send_notification() {
    local message="$1"
    # Slack通知
    curl -X POST -H 'Content-type: application/json' \
         --data "{\"text\":\"不動産スクレーパー: $message\"}" \
         "$SLACK_WEBHOOK_URL" 2>/dev/null || true
         
    # macOS通知
    osascript -e "display notification \"$message\" with title \"不動産スクレーパー\""
}

# 実行
main
```

#### 監視・自動復旧システム
```bash
#!/bin/bash
# monitor_and_recovery.sh

EXPECTED_TIME="02:00"
LOG_FILE="./logs/execution_status.log"

while true; do
    CURRENT_TIME=$(date +%H:%M)
    
    # 実行予定時刻の30分後にチェック
    if [ "$CURRENT_TIME" = "02:30" ]; then
        # 今日のログがあるかチェック
        TODAY=$(date +%Y%m%d)
        if ! grep -q "$TODAY.*スクレイピング開始" "$LOG_FILE"; then
            echo "スクレイピングが実行されていません。手動実行します。"
            ./daily_scraping.sh &
            send_notification "スクレイピング自動復旧実行"
        fi
    fi
    
    # 1時間ごとにチェック
    sleep 3600
done
```

### 🖥️ Option 2: 専用ミニPC（推奨度: 高）

#### Intel NUC / Mac Mini構成
```
Intel NUC 13 Pro:
- CPU: Intel Core i5-1340P
- RAM: 16GB
- SSD: 512GB
- 価格: 約80,000円
- 消費電力: 15W（MacBookより省電力）
- 24/7稼働設計

年間電気代: 15W × 24時間 × 365日 × 30円/kWh = 約3,900円
```

#### 設定例
```bash
# Ubuntu Server セットアップ
# 1. 自動起動設定
sudo systemctl enable scraper-daily.service

# 2. 死活監視
sudo systemctl enable scraper-monitor.service

# 3. 自動再起動（週1回）
echo "0 1 * * 0 /sbin/reboot" | sudo crontab -
```

### 🏠 Option 3: Raspberry Pi（推奨度: 中）

#### Raspberry Pi 4構成
```
Raspberry Pi 4 8GB:
- 価格: 約12,000円
- 消費電力: 6W
- 24/7稼働特化

年間電気代: 6W × 24時間 × 365日 × 30円/kWh = 約1,600円

課題:
- 処理速度がやや遅い（MacBookの1/3程度）
- 画像処理で負荷が高い
- メモリ8GBで制約あり
```

### ☁️ Option 4: ハイブリッド構成（推奨度: 最高）

#### MacBook + VPS バックアップ
```bash
#!/bin/bash
# hybrid_execution.sh

PRIMARY_HOST="localhost"
BACKUP_HOST="your-vps.com"

# プライマリ（MacBook）で実行試行
if ping -c 1 $PRIMARY_HOST > /dev/null 2>&1; then
    ssh user@$PRIMARY_HOST "cd /path/to/scraper && ./daily_scraping.sh"
    if [ $? -eq 0 ]; then
        echo "プライマリ実行成功"
        exit 0
    fi
fi

# プライマリ失敗時はVPSで実行
echo "プライマリ失敗、バックアップVPSで実行"
ssh user@$BACKUP_HOST "cd /path/to/scraper && ./daily_scraping.sh"
```

#### 低価格VPS（バックアップ用）
```
さくらVPS 1G:
- CPU: 1コア
- RAM: 1GB  
- SSD: 50GB
- 価格: 880円/月（年額10,560円）
- 用途: バックアップ実行のみ

Vultr 1GB:
- 価格: $6/月（約840円/月）
- 年額: 約10,000円
```

## 📊 安定性ソリューション比較

| ソリューション | 初期費用 | 年間コスト | 安定性 | 複雑度 | 推奨度 |
|---------------|----------|------------|--------|--------|--------|
| MacBook最適化 | 0円 | 8,000円 | ⭐⭐⭐ | ⭐ | ⭐⭐⭐⭐ |
| 専用ミニPC | 80,000円 | 12,000円 | ⭐⭐⭐⭐⭐ | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| Raspberry Pi | 20,000円 | 8,000円 | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐ |
| ハイブリッド | 0-10,000円 | 18,000円 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |

## 🎯 段階的実装アプローチ

### Phase 1: MacBook安定化（即座に実施）
```bash
# システム設定最適化
sudo pmset -a sleep 0
caffeinate -s &

# 監視スクリプト導入
crontab -e
# 追加: 0 */6 * * * /path/to/monitor_and_recovery.sh
```

### Phase 2: 通知システム構築
```bash
# Slack Webhook設定
export SLACK_WEBHOOK_URL="https://hooks.slack.com/services/YOUR/WEBHOOK/URL"

# 実行状況の外部監視
# UptimeRobot等の無料サービスを活用
```

### Phase 3: バックアップ実行環境（必要に応じて）
- 安価なVPSまたはミニPCを準備
- 自動フェイルオーバー機能実装

## 💡 最終推奨構成

### 🏆 **段階的アプローチ**

#### 今すぐ実施（費用: 0円）
1. MacBookの電源・スリープ設定最適化
2. 監視・通知スクリプト導入
3. 自動復旧機能付きバッチ作成

#### 6ヶ月後検討（予算に余裕ができたら）
1. Intel NUC等の専用ミニPC導入
2. 完全24/7稼働環境構築

#### 安定性の指標
- **現在**: 70-80%（手動管理）
- **Phase 1後**: 90-95%（自動監視）
- **Phase 2後**: 98-99%（専用ハードウェア）

この段階的アプローチにより、コストを抑えながら徐々に安定性を向上できます！