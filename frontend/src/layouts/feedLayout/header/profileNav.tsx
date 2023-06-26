import { useEffect, useState } from 'react';
import { _ROUTES } from '@/helpers/constants';
import { deleteCookie } from 'cookies-next';
import { IconButton, Menu, MenuItem } from '@mui/material';
import { AccountCircle } from '@mui/icons-material';
import { checkAuth } from '@/helpers/checkAuth';
import { useRouter } from 'next/router';

export default function ProfileNav() {
  const router = useRouter();
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(null);
  const [auth, setAuth] = useState(true);
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

  const handleLogout = () => {
    setAuth(false);
    deleteCookie('refresh_token');
    deleteCookie('access_token');
    handleCloseUserMenu();
    router.push(_ROUTES.LADING_PAGE);
  };

  const [isProfileOpen, setIsProfileOpen] = useState(false);
  const handleIsProfileClose = () => {
    setIsProfileOpen(false);
  };
  return (
    <>
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
    </>
  );
}
