import { PlayArrow } from "@mui/icons-material";
import { Button } from "@mui/material";
import React from "react";
import { useNavigate } from "react-router-dom";

const PlayButton = ({ showID }) => {
  const navigate = useNavigate();

  const handlePlayShowEpisode = () => {
    navigate(`/show/${showID}/episode/1`);
  };
  return (
    <Button
      size="large"
      sx={{
        color: "black",
        backgroundColor: "white",
        justifyContent: "center",
        alignContent: "center",
        padding: "5px 10px",
        height: "40px",
        width: "110px",
        marginRight: "7px",
        textTransform: "none",
      }}
      onClick={handlePlayShowEpisode}
    >
      <PlayArrow sx={{ height: "33px", width: "33px", marginRight: "2px" }} />
      <span
        style={{
          fontWeight: "600",
          marginRight: "7.5px",
          fontSize: "16px",
        }}
      >
        Play
      </span>
    </Button>
  );
};

export default PlayButton;
