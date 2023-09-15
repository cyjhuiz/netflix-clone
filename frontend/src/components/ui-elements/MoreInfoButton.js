import { InfoOutlined } from "@mui/icons-material";
import { Button } from "@mui/material";
import React, { useState, Fragment } from "react";
import ShowModal from "../../features/show-modal/components/ShowModal";

const MoreInfoButton = ({ showID }) => {
  const [showModalIsOpen, setShowModalIsOpen] = useState(false);
  const handleOpenShowModal = () => {
    setShowModalIsOpen(true);
  };
  const handleCloseShowModal = () => {
    setShowModalIsOpen(false);
  };
  return (
    <Fragment>
      <Button
        size="large"
        sx={{
          color: "white",
          backgroundColor: "rgba(109, 109, 110, 0.7)",
          justifyContent: "center",
          alignContent: "center",
          padding: "5px 10px",
          height: "40px",
          width: "150px",
          marginRight: "5px",
          textTransform: "none",
        }}
        onClick={handleOpenShowModal}
      >
        <InfoOutlined
          sx={{ height: "25px", width: "25px", marginRight: "4px" }}
        />
        <span
          style={{
            fontWeight: "600",
            marginRight: "7.5px",
            fontSize: "16px",
          }}
        >
          More Info
        </span>
      </Button>
      {showModalIsOpen && (
        <ShowModal
          showModalIsOpen={showModalIsOpen}
          handleCloseShowModal={handleCloseShowModal}
          showID={showID}
        />
      )}
    </Fragment>
  );
};

export default MoreInfoButton;
