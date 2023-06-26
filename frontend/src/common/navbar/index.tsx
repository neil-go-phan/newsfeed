import { _ROUTES } from '@/helpers/constants';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import Menu from '@mui/material/Menu';
import MenuIcon from '@mui/icons-material/Menu';
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import MenuItem from '@mui/material/MenuItem';
import Link from 'next/link';
import React, { useEffect, useState } from 'react';
import { AccountCircle, Feed } from '@mui/icons-material';
import Divider from '@mui/material/Divider';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import Drawer from '@mui/material/Drawer';
import { checkAuth } from '@/helpers/checkAuth';
import { deleteCookie } from 'cookies-next';
import { useRouter } from 'next/router';

const drawerWidth = 240;
type NavPages = {
  link: string;
  name: string;
};
export const pages: Array<NavPages> = [
  { name: 'Features', link: _ROUTES.FEATURE_PAGE },
  { name: 'Pricing', link: _ROUTES.PRICING_PAGE },
  { name: 'Discover', link: _ROUTES.DISCOVER_PAGE },
];

function NavbarComponent() {
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(
    null
  );

  const [mobileOpen, setMobileOpen] = useState(false);
  const [auth, setAuth] = useState(true);
  const router = useRouter();
  useEffect(() => {
    async function checkLogIn() {
      const userChecked: boolean = await checkAuth();
      setAuth(userChecked);
    }

    checkLogIn();
  }, [auth]);

  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const handleDrawerToggle = () => {
    setMobileOpen((prevState) => !prevState);
  };

  const handleLogout = () => {
    setAuth(false);
    deleteCookie('refresh_token');
    deleteCookie('access_token');
    handleCloseUserMenu();
    router.push(_ROUTES.LADING_PAGE);
  };

  const drawer = (
    <Box onClick={handleDrawerToggle} sx={{ textAlign: 'center' }}>
      <Typography variant="h6" sx={{ my: 2 }}>
        Newsfeed
      </Typography>
      <Divider />
      <List>
        {pages.map((page) => (
          <ListItem key={page.name} disablePadding>
            <ListItemButton
              sx={{ my: 1, display: 'block' }}
              className="navbarComponent__mobileNavListItem"
            >
              <Link href={page.link}>{page.name}</Link>
            </ListItemButton>
          </ListItem>
        ))}
      </List>
    </Box>
  );

  return (
    <AppBar position="sticky" className="navbarComponent">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <Feed
            sx={{ display: { xs: 'none', md: 'flex' }, mr: 1, color: 'black' }}
          />
          <Link href={_ROUTES.LADING_PAGE} className="navbarComponent__brand">
            <Typography
              variant="h5"
              noWrap
              component="span"
              sx={{
                mr: 2,
                display: { xs: 'none', md: 'flex' },
                fontFamily: 'monospace',
                fontWeight: 700,
                letterSpacing: '.3rem',
                color: 'black',
              }}
            >
              Newsfeed
            </Typography>
          </Link>

          <Box
            sx={{
              color: 'black',
              flexGrow: 1,
              display: { xs: 'flex', md: 'none' },
            }}
          >
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleDrawerToggle}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Box component="nav">
              <Drawer
                variant="temporary"
                open={mobileOpen}
                onClose={handleDrawerToggle}
                ModalProps={{
                  keepMounted: true,
                }}
                sx={{
                  display: { xs: 'block', sm: 'none' },
                  '& .MuiDrawer-paper': {
                    boxSizing: 'border-box',
                    width: drawerWidth,
                  },
                }}
              >
                {drawer}
              </Drawer>
            </Box>
          </Box>
          <Feed
            sx={{ display: { xs: 'flex', md: 'none' }, mr: 1, color: 'black' }}
          />
          <Link href={_ROUTES.LADING_PAGE} className="navbarComponent__brand--sm">
            <Typography
              variant="h6"
              noWrap
              component="span"
              sx={{
                mr: 2,
                display: { xs: 'flex', md: 'none' },
                fontFamily: 'monospace',
                fontWeight: 700,
                letterSpacing: '.3rem',
                color: 'black',
              }}
            >
              Newsfeed
            </Typography>
          </Link>

          <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
            {pages.map((page) => (
              <Button
                key={page.name}
                sx={{ my: 2, display: 'block' }}
                className="navbarComponent__link mx-3"
              >
                <Link href={page.link}>{page.name}</Link>
              </Button>
            ))}
          </Box>

          <Box sx={{ flexGrow: 0 }} className="navbarComponent__user">
            {auth ? (
              <div className="navbarComponent__user--logged">
                <div className="newsfeedLink">
                  <Link
                    className="mx-md-3 mx-1 btn btn-sm btn-primary d-block px-3 text-nowrap"
                    href={_ROUTES.FEEDS_SEARCH_WEBS}
                  >
                    <span>My newsfeed</span>
                  </Link>
                </div>

                <IconButton
                  size="large"
                  aria-label="account of current user"
                  aria-controls="menu-appbar"
                  aria-haspopup="true"
                  onClick={handleOpenUserMenu}
                  color="inherit"
                >
                  <AccountCircle />
                </IconButton>
                <Menu
                  id="menu-appbar"
                  anchorEl={anchorElUser}
                  anchorOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                  }}
                  keepMounted
                  transformOrigin={{
                    vertical: 'top',
                    horizontal: 'right',
                  }}
                  open={Boolean(anchorElUser)}
                  onClose={handleCloseUserMenu}
                >
                  <MenuItem onClick={handleCloseUserMenu}>Profile</MenuItem>
                  <MenuItem onClick={handleCloseUserMenu}>My account</MenuItem>
                  <MenuItem onClick={handleLogout}>Logout</MenuItem>
                </Menu>
              </div>
            ) : (
              <div className="navbarComponent__user--unlogged">
                <Link className="mx-3 loginBtn" href={_ROUTES.LOGIN_PAGE}>
                  Login
                </Link>
                <Link
                  className="mx-3 btn btn-sm btn-primary d-none d-md-block px-3"
                  href={_ROUTES.REGISTER_PAGE}
                >
                  Create account
                </Link>
              </div>
            )}
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default NavbarComponent;
