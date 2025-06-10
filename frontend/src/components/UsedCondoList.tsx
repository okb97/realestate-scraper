import { useEffect, useState } from "react";

type Condo = {
    ID: number;
    Name: string;
    PriceText: string;
    PriceNum: number;
    Location: string;
    TrainLineStation: string;
    AreaText: string;
    AreaNum: number;
    Layout: string;
    Balcony: string;
    BuiltAt: string;     // `time.Time` は ISO8601 文字列で来るので string で受ける
    URL: string;
    ScrapedAt: string;
    CreatedAt: string;
    UpdatedAt: string;
    IsActive: boolean;
  };

export default function UsedCondoList() {
    const [condos, setCondos] = useState<Condo[]>([]);
    const [loading, setLoading] = useState(true);
  
    useEffect(() => {
      fetch("http://localhost:8080/api/used-condos")
        .then((res) => res.json())
        .then((data) => {
            // console.log("取得したデータ:", data);
          setCondos(data);
          setLoading(false);
        })
        .catch((error) => {
          console.error("API取得エラー:", error);
          setLoading(false);
        });
    }, []);
  
    if (loading) return <p>読み込み中...</p>;
  
    return (
        <div style={{ backgroundColor: "white", color: "black", minHeight: "100vh", padding: "20px" }}>
          <h2>中古マンション一覧</h2>
          <ul style={{ padding: 0, listStyle: "none" }}>
            {condos.map((condo) => (
              <li
                key={condo.ID}
                style={{
                  border: "1px solid #ccc",
                  borderRadius: "8px",
                  padding: "12px",
                  marginBottom: "12px",
                  backgroundColor: "#f9f9f9",
                  boxShadow: "0 1px 3px rgba(0,0,0,0.1)",
                }}
              >
                <strong style={{ fontSize: "1.2em" }}>{condo.Name}</strong>
                <p>所在地: {condo.Location}</p>
                <p>沿線・駅: {condo.TrainLineStation}</p>
                <p>間取り: {condo.Layout}</p>
                <p>面積: {condo.AreaText} ({condo.AreaNum}㎡)</p>
                <p>価格: {condo.PriceText}</p>
                <p>築年: {new Date(condo.BuiltAt).toLocaleDateString()}</p>
                <a
                  href={condo.URL}
                  target="_blank"
                  rel="noopener noreferrer"
                  style={{
                    display: "inline-block",
                    marginTop: "8px",
                    padding: "6px 12px",
                    backgroundColor: "#007bff",
                    color: "white",
                    borderRadius: "4px",
                    textDecoration: "none",
                    fontWeight: "bold",
                  }}
                >
                  詳細を見る
                </a>
              </li>
            ))}
          </ul>
        </div>
      );
  }
  