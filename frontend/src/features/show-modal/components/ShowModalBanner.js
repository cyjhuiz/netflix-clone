import { Add, PlayArrow, ThumbUpOffAlt } from "@mui/icons-material";
import { Button, IconButton } from "@mui/material";
import React, { useState } from "react";
import PlayButton from "../../../components/ui-elements/PlayButton";
import FavouriteButton from "../../../components/ui-elements/FavouriteButton";
import LikeButton from "../../../components/ui-elements/LikeButton";

const ShowModalBanner = ({ show }) => {
  const hasEpisodes = !!show?.episodes;
  console.log(hasEpisodes);
  return (
    <header
      className="banner"
      style={{
        backgroundSize: "cover",
        backgroundImage: `url(
            ${show?.thumbnailURL}
        )`,
        backgroundPosition: "center center",
        objectFit: "contain",
        display: "flex",
        height: "450px",
      }}
    >
      <div
        style={{
          height: "auto",
          width: "63%",
          display: "flex",
          flexDirection: "column-reverse",
          color: "blue",
          padding: "40px 40px",
        }}
      >
        <div style={{ display: "flex", flexDirection: "row" }}>
          <PlayButton showID={show?.showID} disabled={!hasEpisodes} />
          <FavouriteButton showID={show?.showID} />
          <LikeButton showID={show?.showID} />
        </div>
      </div>
    </header>
  );
};

export default ShowModalBanner;
