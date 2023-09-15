import { NotificationsNone } from "@mui/icons-material";
import { IconButton } from "@mui/material";
import React, { useContext, useEffect, useState } from "react";
import PopupState, { bindHover } from "material-ui-popup-state";
import NavNotificationPopover from "./NavNotificationPopover";
import { AuthContext } from "../../features/auth/context/auth-context";
import { BACKEND_NOTIFICATIONS_API_BASE_URL } from "../../util/constants";
import axios from "axios";

const NavNotificationIconButton = () => {
  const { userID } = useContext(AuthContext);
  const [notifications, setNotifications] = useState([]);

  useEffect(() => {
    const fetchNotifications = async () => {
      const response = await axios
        .get(
          `${BACKEND_NOTIFICATIONS_API_BASE_URL}/notification?userID=${userID}`
        )
        .catch((err) => console.log(err));

      if (response?.status === 200) {
        console.log(response.status);
        setNotifications(response.data);
      }
    };

    if (!!userID) {
      fetchNotifications();
    }
  }, [userID]);

  return (
    <PopupState variant="popover" popupId="notification-popover">
      {(popupState) => (
        <div>
          <IconButton
            {...bindHover(popupState)}
            aria-haspopup="true"
            sx={{ p: 0 }}
            disableRipple
          >
            <NotificationsNone
              sx={{ height: "28px", width: "28px", color: "#ffffffd6" }}
            />
          </IconButton>
          <NavNotificationPopover
            notifications={notifications}
            popupState={popupState}
          />
        </div>
      )}
    </PopupState>
  );
};

export default NavNotificationIconButton;
