import { ThumbUpAlt, ThumbUpOffAlt } from "@mui/icons-material";
import { IconButton, Tooltip } from "@mui/material";
import React, { useCallback, useContext, useEffect, useState } from "react";
import { AuthContext } from "../../features/auth/context/auth-context";
import { BACKEND_SHOW_API_BASE_URL } from "../../util/constants";
import axios from "axios";

const LikeButton = ({ showID }) => {
  const { userID } = useContext(AuthContext);

  const [hasLiked, setHasLiked] = useState(false);

  const fetchLike = useCallback(async (userID, showID) => {
    const response = await axios
      .get(`${BACKEND_SHOW_API_BASE_URL}/show/${showID}/like?userID=${userID}`)
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasLiked(!!response.data);
    }
  }, []);

  const handleLike = async () => {
    const response = await axios
      .post(`${BACKEND_SHOW_API_BASE_URL}/show/${showID}/like?userID=${userID}`)
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasLiked(true);
    }
  };

  const handleUnlike = async () => {
    const response = await axios
      .delete(
        `${BACKEND_SHOW_API_BASE_URL}/show/${showID}/like?userID=${userID}`
      )
      .catch((err) => console.log(err));

    if (response?.status === 200) {
      setHasLiked(false);
    }
  };

  useEffect(() => {
    fetchLike(userID, showID);
  }, [fetchLike, userID, showID]);

  return !hasLiked ? (
    <IconButton
      onClick={handleLike}
      sx={{
        p: 0,
        border: "1px solid rgba(255, 255, 255, 0.7)",
        backgroundColor: "rgba(42,42,42,.6)",
        padding: "7px",
      }}
      disableRipple
    >
      <Tooltip title="Add to likes" placement="top">
        <ThumbUpOffAlt sx={{ height: "24px", width: "24px", color: "white" }} />
      </Tooltip>
    </IconButton>
  ) : (
    <IconButton
      onClick={handleUnlike}
      sx={{
        p: 0,
        border: "1px solid rgba(255, 255, 255, 0.7)",
        backgroundColor: "rgba(42,42,42,.6)",
        padding: "7px",
      }}
      disableRipple
    >
      <Tooltip title="Remove from likes" placement="top">
        <ThumbUpAlt sx={{ height: "24px", width: "24px", color: "white" }} />
      </Tooltip>
    </IconButton>
  );
};

export default LikeButton;
