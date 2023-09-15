import React, { useEffect, useState } from "react";
import "./FeaturedShowBanner.css";
import axios from "axios";
import { BACKEND_SHOW_API_BASE_URL } from "../../../util/constants";
import PlayButton from "../../../components/ui-elements/PlayButton";
import MoreInfoButton from "../../../components/ui-elements/MoreInfoButton";

const truncateStr = (str, num) =>
  str?.length > num ? `${str.substring(0, num)}...` : str;
const FeaturedShowBanner = () => {
  const [show, setShow] = useState();

  useEffect(() => {
    const fetchFeaturedShow = async () => {
      const featuredCategory = "Netflix Originals";

      const response = await axios
        .get(`${BACKEND_SHOW_API_BASE_URL}/show?category=${featuredCategory}`)
        .catch((err) => console.log("Failed to fetch featured movie"));

      if (response?.status === 200) {
        const shows = response.data;
        const randomShowIdx = Math.floor(Math.random() * (shows?.length - 1));
        const featuredShow = shows[randomShowIdx];
        setShow(featuredShow);
      }
    };

    fetchFeaturedShow();
  }, []);
  return (
    <header
      className="banner"
      style={{
        backgroundSize: "cover",
        backgroundImage: `url(
            ${show?.thumbnailURL}
        )`,
        backgroundPosition: "center center",
      }}
    >
      <div className="banner__contents">
        <h1 className="banner__title">{show?.title}</h1>
        <div className="banner__buttons">
          <PlayButton showID={show?.showID} />
          <MoreInfoButton showID={show?.showID} />
        </div>

        <h1 className="banner__description">
          {truncateStr(show?.description, 150)}
        </h1>
      </div>
      <div className="banner__fadeBottom"></div>
    </header>
  );
};

export default FeaturedShowBanner;
