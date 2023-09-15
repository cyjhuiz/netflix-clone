import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import { IconButton } from "@mui/material";
export default function LandingPageNavBar({ setMode }) {
  return (
    <Box sx={{ flexGrow: 1, border: "none" }}>
      <AppBar color="transparent" position="static" elevation={0}>
        <Toolbar sx={{ justifyContent: "space-between" }}>
          <IconButton>
            <img
              src="/images/logo.png"
              href="/test"
              alt="netflix logo"
              style={{
                height: "80px",
                width: "auto",
                cursor: "pointer",
                border: "none",
              }}
            />
          </IconButton>

          <Button
            disableRipple
            sx={{
              backgroundColor: "#E50914",
              color: "white",
              fontWeight: "600",
              borderRadius: 0,
              fontSize: "0.rem",
              padding: "8px 20px",
              "&:hover": {
                background: "#E50914",
              },
            }}
            onClick={() => setMode("LOGIN")}
          >
            Sign In
          </Button>
        </Toolbar>
      </AppBar>
    </Box>
  );
}
