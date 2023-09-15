import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import Container from "@mui/material/Container";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import MenuItem from "@mui/material/MenuItem";
import { Search } from "@mui/icons-material";
import { AuthContext } from "../../features/auth/context/auth-context";
import NavNotificationIconButton from "./NavNotificationIconButton";

const pages = [
  "Home",
  "Series",
  "Films",
  "New & Popular",
  "My List",
  "Browse By Languages",
];

function MainNavBar() {
  const { logout } = React.useContext(AuthContext);

  const [anchorElNav, setAnchorElNav] = React.useState(null);
  const [anchorElUser, setAnchorElUser] = React.useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const handleLogout = () => {
    logout();
    handleCloseNavMenu();
  };

  const [isShown, setIsShown] = React.useState(false);

  React.useEffect(() => {
    window.addEventListener("scroll", () => {
      if (window.scrollY > 100) {
        setIsShown(true);
      } else {
        setIsShown(false);
      }
    });

    return () => {
      window.removeEventListener("scroll", null);
    };
  }, []);

  return (
    <AppBar
      position="fixed"
      elevation={0}
      sx={{
        zIndex: "1",
        height: "63px",
        backgroundColor: `${isShown ? "#111" : "transparent"}`,
        transitionTimingFunction: "ease-in",
        transition: "all 0.5s",
      }}
    >
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <IconButton disableRipple>
            <img
              src="/images/logo.png"
              href="/test"
              alt="netflix logo"
              style={{
                height: "50px",
                width: "auto",
                cursor: "pointer",
                border: "none",
              }}
            />
          </IconButton>

          <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: "block", md: "none" },
              }}
            >
              {pages.map((page) => (
                <MenuItem key={page} onClick={handleCloseNavMenu}>
                  <Typography textAlign="center">{page}</Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>

          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            {pages.map((page) => (
              <Button
                key={page}
                onClick={handleCloseNavMenu}
                sx={{
                  my: 2,
                  color: "#ffffffd6",
                  display: "block",
                  textTransform: "none",
                  "&:hover": {
                    background: "transparent",
                  },
                }}
              >
                <Typography sx={{ fontWeight: "500" }}>{page}</Typography>
              </Button>
            ))}
          </Box>

          <Box
            sx={{
              flexGrow: 0,
              display: "flex",
              columnGap: "15px",
              justifyContent: "space-between",
            }}
          >
            <IconButton sx={{ p: 0 }} disableRipple>
              <Search
                sx={{ height: "28px", width: "28px", color: "#ffffffd6" }}
              />
            </IconButton>
            <Button
              sx={{
                p: 0,
                color: "#ffffffd6",
                textTransform: "none",
                "&:hover": {
                  background: "transparent",
                },
              }}
              disableRipple
            >
              Children
            </Button>
            <NavNotificationIconButton />
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                <Avatar
                  variant="square"
                  alt="user avatar icon"
                  src="/images/default-slate.png"
                  sx={{ height: "32px", width: "32px" }}
                />
              </IconButton>
            </Tooltip>
            <Menu
              sx={{ mt: "45px" }}
              id="menu-appbar"
              anchorEl={anchorElUser}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
              disableScrollLock={true}
            >
              <MenuItem key={"menuItem_profile"} onClick={handleCloseUserMenu}>
                <Typography textAlign="center">Profile</Typography>
              </MenuItem>
              <MenuItem key={"menuItem_account"} onClick={handleCloseUserMenu}>
                <Typography textAlign="center">Account</Typography>
              </MenuItem>
              <MenuItem key="menuItem_logout" onClick={handleLogout}>
                <Typography textAlign="center">Logout</Typography>
              </MenuItem>
            </Menu>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default MainNavBar;
