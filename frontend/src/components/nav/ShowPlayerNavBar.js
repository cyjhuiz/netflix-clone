import { KeyboardBackspace } from "@mui/icons-material";
import { AppBar, Container, IconButton, Toolbar } from "@mui/material";
import React from "react";
import { useNavigate } from "react-router-dom";

const ShowPlayerNavBar = ({ navbarIsVisible }) => {
  const navigate = useNavigate();
  const handleGoBack = () => {
    navigate("/browse");
  };
  return (
    <AppBar
      position="fixed"
      elevation={0}
      sx={{
        zIndex: "1",
        height: "63px",
        backgroundColor: "transparent",
        transitionTimingFunction: "ease-in",
        transition: "all 0.5s",
        visibility: navbarIsVisible ? "visible" : "hidden",
      }}
    >
      <Container maxWidth={false} disableGutters style={{ paddingLeft: "4px" }}>
        <Toolbar disableGutters>
          <IconButton disableRipple onClick={handleGoBack}>
            <KeyboardBackspace
              sx={{
                color: "white",
                height: "50px",
                width: "50px",
                marginTop: "36px",
              }}
            />
          </IconButton>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default ShowPlayerNavBar;
