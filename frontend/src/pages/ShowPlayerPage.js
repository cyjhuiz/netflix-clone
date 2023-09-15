import axios from "axios";
import React, { Fragment, useEffect, useState } from "react";
import ReactPlayer from "react-player/youtube";
import { BACKEND_SHOW_API_BASE_URL } from "../util/constants";
import { useParams } from "react-router-dom";
import ShowPlayerNavBar from "../components/nav/ShowPlayerNavBar";
import movieTrailer from "movie-trailer";
import { Snackbar } from "@mui/material";

let navbarTimeout;

const ShowPlayerPage = () => {
  const { showID, number } = useParams();
  const [episode, setEpisode] = useState(undefined);
  const [placeholderVideoURL, setPlaceholderVideoURL] = useState(undefined);
  const [navbarIsVisible, setNavbarIsVisible] = useState(true);

  const [snackbar, setSnackbar] = React.useState({
    isOpen: false,
    message: "",
  });

  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }

    setSnackbar((prev) => ({
      ...snackbar,
      isOpen: false,
      message: "",
    }));
  };

  useEffect(() => {
    const fetchEpisode = async () => {
      const response = await axios
        .get(`${BACKEND_SHOW_API_BASE_URL}/show/${showID}/episode/${number}`)
        .catch((err) => {
          console.log(err);

          setSnackbar((prev) => ({
            ...snackbar,
            isOpen: true,
            message: `No relevant episodes found. Loading placeholder video.`,
          }));
        });

      if (response?.status === 200) {
        setEpisode(response.data);
      }
    };

    fetchEpisode();
  }, [showID, number, snackbar]);

  useEffect(() => {
    const fetchShow = async () => {
      const response = await axios
        .get(`${BACKEND_SHOW_API_BASE_URL}/show/${showID}`)
        .catch((err) => console.log(err));

      if (response?.status === 200) {
        const show = response.data;

        let videoURL = "https://www.youtube.com/watch?v=u31qwQUeGuM";
        const searchedVideoURL = await movieTrailer(show.title);
        if (searchedVideoURL !== null) {
          videoURL = searchedVideoURL;
        }

        setPlaceholderVideoURL(videoURL);
      }
    };

    fetchShow();
  }, [showID]);

  const handleMouseMove = () => {
    clearTimeout(navbarTimeout);
    setNavbarIsVisible(true);

    navbarTimeout = setTimeout(() => {
      setNavbarIsVisible(false);
    }, 2600);
  };

  return (
    <div
      style={{
        height: "100vh",
        minHeight: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        overflow: "hidden",
      }}
    >
      <Snackbar
        open={snackbar.isOpen}
        autoHideDuration={6000}
        onClose={handleClose}
        message={snackbar.message}
      />
      {!!episode || !!placeholderVideoURL ? (
        <Fragment>
          <ShowPlayerNavBar navbarIsVisible={navbarIsVisible} />
          <ReactPlayer
            url={episode?.videoURL || placeholderVideoURL}
            width="100%"
            height="100%"
            controls={true}
            onMouseMove={handleMouseMove}
          />
        </Fragment>
      ) : (
        <Fragment />
      )}
    </div>
  );
};

export default ShowPlayerPage;
