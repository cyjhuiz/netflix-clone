import React, { useState } from "react";
import "./ShowCard.css";
import ShowModal from "../../show-modal/components/ShowModal";

const ShowCard = ({ show }) => {
  const [showModalIsOpen, setShowModalIsOpen] = useState(false);
  const handleOpenShowModal = () => {
    setShowModalIsOpen(true);
  };
  const handleCloseShowModal = () => {
    setShowModalIsOpen(false);
  };
  return (
    <>
      <img
        className={`showCard`}
        src={`${show.thumbnailURL}`}
        alt={show.title}
        onClick={handleOpenShowModal}
      />
      {showModalIsOpen && (
        <ShowModal
          showModalIsOpen={showModalIsOpen}
          handleOpenShowModal={handleOpenShowModal}
          handleCloseShowModal={handleCloseShowModal}
          showID={show.showID}
        />
      )}
    </>
  );
};

export default ShowCard;
