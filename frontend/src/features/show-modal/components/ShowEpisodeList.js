import React from "react";
import ShowEpisodeCard from "./ShowEpisodeCard";

const ShowEpisodeList = ({ episodes }) => {
  return (
    <>
      <div>
        <h2 style={{ marginTop: "0" }}>Episodes</h2>
      </div>
      <div>
        {episodes?.map((episode, index) => (
          <ShowEpisodeCard key={`showEpisodeCard_${index}`} episode={episode} />
        ))}
      </div>
    </>
  );
};

export default ShowEpisodeList;
