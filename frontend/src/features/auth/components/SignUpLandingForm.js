import { Box, Button } from "@mui/material";
import React from "react";

const SignUpLandingForm = ({ email, setEmail, setMode }) => {
  return (
    <Box
      sx={{
        color: "white",
        position: "absolute",
        top: "50%",
        left: "50%",
        marginRight: "-50%",
        transform: "translate(-50%, -50%)",
        textAlign: "center",
      }}
    >
      <h1 style={{ fontSize: "50px" }}>
        Unlimited films, TV programmes and more
      </h1>
      <h3 style={{ fontSize: "30px" }}>Watch anywhere. Cancel at any time</h3>
      <h4 style={{ fontSize: "20px" }}>
        Ready to watch? Enter your email to create or restart your membership.
      </h4>
      <Box>
        <input
          label="Email"
          type="email"
          placeholder="Email"
          style={{
            backgroundColor: "white",
            height: "42.5px",
            width: "30%",
            maxWidth: "600px",
            outline: "none",
            borderWidth: "0",
            "*:focus": {
              outline: "none",
            },
          }}
          onChange={(event) => setEmail(event.target.value)}
          value={email}
        />
        <Button
          sx={{
            backgroundColor: "#E50914",
            color: "white",
            borderRadius: 0,
            height: "45px",
            fontSize: "1 rem",
            padding: "8px 20px",
            "&:hover": {
              background: "#E50914",
            },
          }}
          onClick={() => setMode("SIGN_UP")}
        >
          Get Started
        </Button>
      </Box>
    </Box>
  );
};

export default SignUpLandingForm;
