import { Box, Modal } from "@mui/material";
import React, { Fragment, useEffect, useState } from "react";
import ShowInfo from "./ShowInfo";
import ShowMetaInfo from "./ShowMetaInfo";
import ShowModalBanner from "./ShowModalBanner";
import "./ShowModal.css";
import ShowEpisodeList from "./ShowEpisodeList";
import axios from "axios";
import { BACKEND_SHOW_API_BASE_URL } from "../../../util/constants";

const ShowModal = ({ showModalIsOpen, handleCloseShowModal, showID }) => {
  console.log("showID", showID);
  const [show, setShow] = useState();
  useEffect(() => {
    const fetchShow = async () => {
      console.log("fetchingshow showID:", showID);
      const response = await axios
        .get(`${BACKEND_SHOW_API_BASE_URL}/show/${showID}`)
        .catch((err) => console.log(err));
      if (response?.status === 200) {
        setShow(response.data);
      }
    };

    if (showModalIsOpen && !!showID) {
      fetchShow();
    }
  }, [showModalIsOpen, showID]);

  return (
    <Modal
      open={showModalIsOpen}
      onClose={handleCloseShowModal}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
      sx={{
        justifyContent: "center",
        display: "flex",
        overflow: "hidden",
      }}
    >
      {!!show ? (
        <Box
          className="showModal"
          style={{
            backgroundColor: "#111",
            width: "63%",
            borderRadius: "10px",
            borderColor: "transparent",
            overflow: "hidden",
            color: "white",
            marginTop: "4rem",
            overflowY: "scroll",
            scrollbars: {
              display: "none",
            },
          }}
        >
          <ShowModalBanner show={show} />
          <div style={{ padding: "25px 50px" }}>
            <div style={{ display: "flex", flexDirection: "row" }}>
              <div style={{ width: "65%" }}>
                <ShowInfo show={show} />
              </div>
              <div style={{ marginLeft: "10px", width: "35%" }}>
                <ShowMetaInfo />
              </div>
            </div>
            <div style={{ marginTop: !!show.description ? "10px" : "0px" }}>
              <ShowEpisodeList episodes={show.episodes} />
            </div>
          </div>
        </Box>
      ) : (
        <Fragment />
      )}
    </Modal>
  );
};

export default ShowModal;
