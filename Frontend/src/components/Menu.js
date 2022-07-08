import React from 'react';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import { Link , useNavigate } from 'react-router-dom';
import { useStateValue } from "../StateProvider";
import { actionTypes } from "../reducer";
import { makeStyles } from "@material-ui/core/styles";
import { Typography } from '@material-ui/core';
import MenuIcon from '@material-ui/icons/Menu';

const useStyles = makeStyles((theme) => ({
    link: {
      textDecoration: "none",
    },
  }));

export default function SimpleMenu() {
  const classes = useStyles();
  const [anchorEl, setAnchorEl] = React.useState(null);
  const [{ user }, dispatch] = useStateValue();
  const navigate = useNavigate();

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleAuth = () => {
    if (user) {
      dispatch({
        type: actionTypes.SET_USER,
        user: null,
      });
      dispatch({
        type: actionTypes.EMPTY_BASKET,
        basket: [],
      });
      navigate('/');
    }
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  return (
    <div>
      <Button aria-controls="simple-menu" aria-haspopup="true" onClick={handleClick}>
       <MenuIcon fontSize='large'></MenuIcon>
      </Button>
      <Menu
        id="simple-menu"
        anchorEl={anchorEl}
        keepMounted
        open={Boolean(anchorEl)}
        onClose={handleClose}
      >
        <MenuItem onClick={handleClose}><Link to="/orders" className={classes.link}><Button><Typography  color="textPrimary">my orders</Typography></Button></Link></MenuItem>
        <MenuItem onClick={handleClose && handleAuth}><Button><Typography  color="textPrimary">logout</Typography></Button></MenuItem>
      </Menu>
    </div>
  );
}