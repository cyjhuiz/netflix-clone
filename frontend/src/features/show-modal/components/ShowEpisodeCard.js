import React from "react";
import { useNavigate } from "react-router-dom";

const ShowEpisodeCard = ({ episode }) => {
  const navigate = useNavigate();
  const handlePlayEpisode = () => {
    navigate(`/show/${episode.showID}/episode/${episode.number}`);
  };
  return (
    <div
      style={{
        display: "flex",
        flexdirection: "row",
        padding: "15px 15px",
        alignItems: "center",
        justifyContent: "center",
        borderTop: "1px solid #404040",
        borderRadius: "4px",
        cursor: "pointer",
      }}
      onClick={handlePlayEpisode}
    >
      <h2 style={{ margin: "0 15px" }}>{episode.number}</h2>
      <img
        style={{
          borderRadius: "5px",
          height: "6rem",
          width: "auto",
        }}
        src={episode.thumbnailURL}
        alt="episode thumbnail"
      />
      <div style={{ padding: "10px" }}>
        <h3 style={{ margin: "0" }}>{episode.title}</h3>
        <span>{episode.description}</span>
      </div>
    </div>
  );
};

export default ShowEpisodeCard;
