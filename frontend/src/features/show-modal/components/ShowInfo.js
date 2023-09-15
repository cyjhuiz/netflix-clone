import React from "react";

const ShowInfo = ({ show }) => {
  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <h1 style={{ margin: "0" }}>{show.title}</h1>
      <span style={{ paddingTop: "10px" }}>{show?.description}</span>
    </div>
  );
};

export default ShowInfo;
