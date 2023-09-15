import { bindPopover } from "material-ui-popup-state";
import HoverPopover from "material-ui-popup-state/HoverPopover";
import React from "react";

const NavNotificationPopover = ({ popupState, notifications }) => {
  return (
    <HoverPopover
      {...bindPopover(popupState)}
      disableScrollLock={true}
      anchorOrigin={{
        vertical: "bottom",
        horizontal: "right",
      }}
      transformOrigin={{
        vertical: "top",
        horizontal: "right",
      }}
    >
      <div
        style={{
          backgroundColor: "#111",
          color: "gray",
          height: "auto",
          width: "340px",
          fontSize: "14px",
          overflowY: "scroll",
          overflow: "hidden",
        }}
        sx={{
          pointerEvents: "auto",
        }}
      >
        <div style={{ width: "auto", border: "0.5px solid white" }}></div>
        {notifications?.length > 0 ? (
          notifications?.map((notification, index) => (
            <div
              key={`notification_${index}`}
              style={{
                height: "70px",
                display: "flex",
                padding: "10px 7px 10px 12.5px",
                borderTop: "1px solid gray",
                alignItems: "center",
                backgroundColor: "inherit",
              }}
            >
              <img
                src={notification.thumbnailURL}
                alt="notification thumbnail"
                style={{
                  borderRadius: "5px",
                  height: "60px",
                  width: "auto",
                  marginRight: "20px",
                }}
              />
              <div
                style={{
                  display: "flex",
                  flexDirection: "column",
                  alignItems: "start",
                  alignSelf: "start",
                }}
              >
                <span>{notification.title}</span>
                <span>{notification.description}</span>
              </div>
            </div>
          ))
        ) : (
          <div
            style={{
              height: "40px",
              display: "flex",
              padding: "10px 7px 10px 12.5px",
              borderTop: "1px solid gray",
              alignItems: "center",
              justifyContent: "center",
              backgroundColor: "inherit",
            }}
          >
            <span>No notifications to display.</span>
          </div>
        )}
      </div>
    </HoverPopover>
  );
};

export default NavNotificationPopover;
