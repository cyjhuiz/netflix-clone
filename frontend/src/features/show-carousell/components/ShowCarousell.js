import React, { useState, useEffect } from "react";
import ShowCard from "./ShowCard";
import "./ShowCarousell.css";
import axios from "axios";
import { BACKEND_SHOW_API_BASE_URL } from "../../../util/constants";

const ShowCarousell = ({ category }) => {
  const [shows, setShows] = useState([]);

  useEffect(() => {
    const fetchMovies = async () => {
      const response = await axios
        .get(`${BACKEND_SHOW_API_BASE_URL}/show?category=${category}`)
        .catch((err) => {
          console.log("Failed to load movies");
          return;
        });

      console.log(response);
      if (response?.status === 200) {
        const shows = response.data;
        setShows(shows);
        console.log(shows);
        return;
      }
    };

    fetchMovies();
  }, [category]);

  return (
    <div>
      <h2 className="showCarousell">{category}</h2>
      <div className="showCarousell__showCards">
        {shows?.map((show, index) => (
          <ShowCard key={`showCard_${index}`} category={category} show={show} />
        ))}
      </div>
    </div>
  );
};

export default ShowCarousell;
