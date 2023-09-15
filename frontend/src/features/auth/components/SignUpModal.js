import { TextField, Box, Button, Snackbar } from "@mui/material";
import axios from "axios";
import React, { Fragment, useState } from "react";
import { BACKEND_USER_API_BASE_URL } from "../../../util/constants";

const SignUpModal = ({ setMode, initialInputEmail }) => {
  const [email, setEmail] = useState(initialInputEmail);
  const [password, setPassword] = useState("");

  const [snackbar, setSnackbar] = React.useState({
    isOpen: false,
    message: "",
  });

  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }

    setSnackbar((prev) => ({
      ...snackbar,
      isOpen: false,
      message: "",
    }));
  };

  const handleSignUp = () => {
    console.log("signing up");
    const sendSignUpRequest = async () => {
      const signUpRequest = { email, password };
      const response = await axios
        .post(`${BACKEND_USER_API_BASE_URL}/user`, signUpRequest)
        .catch((err) => {
          console.log(err);
          setSnackbar((prev) => ({
            ...snackbar,
            isOpen: true,
            message: err.response.data.error,
          }));
        });
      console.log(response);
      if (response?.status === 200) {
        setSnackbar((prev) => ({
          ...snackbar,
          isOpen: true,
          message: "Sign up success! Redirecting to sign in page",
        }));

        setTimeout(() => {
          setMode("LOGIN");
        }, 3000);
      }
    };

    sendSignUpRequest();
  };
  return (
    <Fragment>
      <Snackbar
        open={snackbar.isOpen}
        autoHideDuration={6000}
        onClose={handleClose}
        message={snackbar.message}
      />
      <Box
        sx={{
          backgroundColor: "rgba(0, 0, 0, 0.85)",
          color: "white",
          position: "absolute",
          top: "55%",
          left: "50%",
          marginRight: "-50%",
          transform: "translate(-50%, -50%)",
          minWidth: "450px",
          minHeight: "450px",
          maxWidth: "300px",
          textAlign: "center",
        }}
      >
        <div
          style={{
            margin: "16%",
            minWidth: "250px",
          }}
        >
          <h1 style={{ textAlign: "left" }}>Sign up</h1>
          <TextField
            id="outlined-basic"
            label="Email"
            variant="outlined"
            fullWidth
            sx={{
              backgroundColor: "white",
              borderRadius: "5px",
              "& label.Mui-focused": {
                color: "gray",
              },
              "& .MuiOutlinedInput-root": {
                "&.Mui-focused fieldset": {
                  borderColor: "gray",
                },
              },
            }}
            onChange={(event) => setEmail(event.target.value)}
            value={email}
          />
          <TextField
            id="outlined-basic"
            label="Password"
            type="password"
            variant="outlined"
            fullWidth
            sx={{
              backgroundColor: "white",
              marginTop: "12.5px",
              borderRadius: "5px",
              "& label.Mui-focused": {
                color: "gray",
              },
              "& .MuiOutlinedInput-root": {
                "&.Mui-focused fieldset": {
                  borderColor: "gray",
                },
              },
            }}
            onChange={(event) => setPassword(event.target.value)}
            value={password}
          />
          <br />
          <Button
            fullWidth
            sx={{
              backgroundColor: "#E50914",
              color: "white",
              height: "50px",
              fontSize: "1 rem",
              fontWeight: "600",
              padding: "8px 20px",
              "&:hover": {
                background: "#E50914",
              },
              marginTop: "25px",
              borderRadius: "5px",
            }}
            onClick={handleSignUp}
          >
            Sign up
          </Button>
          <h4 style={{ textAlign: "left", marginTop: "15px" }}>
            <span style={{ color: "gray" }}>Have an account? </span>
            <span
              style={{ cursor: "pointer" }}
              onClick={() => setMode("LOGIN")}
            >
              Sign in now.
            </span>
          </h4>
        </div>
      </Box>
    </Fragment>
  );
};

export default SignUpModal;
