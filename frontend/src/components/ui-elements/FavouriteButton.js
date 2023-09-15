import { Add, DoneOutlined } from "@mui/icons-material";
import { IconButton, Tooltip } from "@mui/material";
import axios from "axios";
import React, { useCallback, useContext, useEffect, useState } from "react";
import { BACKEND_SHOW_API_BASE_URL } from "../../util/constants";
import { AuthContext } from "../../features/auth/context/auth-context";

const FavouriteButton = ({ showID }) => {
  const { userID } = useContext(AuthContext);

  const [hasFavourited, setHasFavourited] = useState(false);

  const fetchFavourite = useCallback(async (userID, showID) => {
    const response = await axios
      .get(
        `${BACKEND_SHOW_API_BASE_URL}/show/${showID}/favourite?userID=${userID}`
      )
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasFavourited(!!response.data);
    }
  }, []);

  const handleAddToFavourites = async () => {
    const response = await axios
      .post(
        `${BACKEND_SHOW_API_BASE_URL}/show/${showID}/favourite?userID=${userID}`
      )
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasFavourited(true);
    }
  };

  const handleDeleteFavourite = async () => {
    const response = await axios
      .delete(
        `${BACKEND_SHOW_API_BASE_URL}/show/${showID}/favourite?userID=${userID}`
      )
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasFavourited(false);
    }
  };

  useEffect(() => {
    fetchFavourite(userID, showID);
  }, [fetchFavourite, userID, showID]);

  return !hasFavourited ? (
    <IconButton
      onClick={handleAddToFavourites}
      sx={{
        p: 0,
        border: "1px solid rgba(255, 255, 255, 0.7)",
        backgroundColor: "rgba(42,42,42,.6)",
        padding: "6px",
        marginRight: "3px",
      }}
      disableRipple
    >
      <Tooltip title="Add to Favourites" placement="top">
        <Add sx={{ height: "26px", width: "26px", color: "#ffffffd6" }} />
      </Tooltip>
    </IconButton>
  ) : (
    <IconButton
      onClick={handleDeleteFavourite}
      sx={{
        p: 0,
        border: "1px solid rgba(255, 255, 255, 0.7)",
        backgroundColor: "rgba(42,42,42,.6)",
        padding: "6px",
        marginRight: "3px",
      }}
      disableRipple
    >
      <Tooltip title="Remove from Favourites" placement="top">
        <DoneOutlined
          sx={{ height: "26px", width: "26px", color: "#ffffffd6" }}
        />
      </Tooltip>
    </IconButton>
  );
};

export default FavouriteButton;
