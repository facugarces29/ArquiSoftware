import { alpha, makeStyles } from '@material-ui/core/styles';
import { React , useState } from "react";
import InputBase from '@material-ui/core/InputBase';
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import { Badge, Button, CssBaseline } from "@material-ui/core";
import { ShoppingCart } from "@material-ui/icons";
import { useStateValue } from "../StateProvider";
import { Link, useNavigate } from "react-router-dom";
import SearchIcon from '@material-ui/icons/Search';
import { actionTypes , getItemsQuantity } from "../reducer";
import Menu from "./Menu"

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    marginBottom: "7rem",
  },
  appBar: {
    backgroundColor:
      theme.palette.type === 'light' ? theme.palette.grey[200] : theme.palette.grey[800],
    boxShadow: "none",
  },
  grow: {
    flexGrow: 1,
  },
  button: {
    marginLeft: theme.spacing(2),
  },
  image: {
    marginRight: "10px",
  },
  link: {
    textDecoration: "none",
  },
  search: {
    position: 'relative',
    borderRadius: theme.shape.borderRadius,
    borderWidth: '1px',
    borderStyle: 'solid',
    borderColor: 'lightgrey',
    backgroundColor: alpha(theme.palette.common.white, 0.15),
    '&:hover': {
      backgroundColor: alpha(theme.palette.common.white, 0.25),
    },
    marginLeft: 1,
    marginRight: 2,
    width: '20%',
  },
  searchIcon: {
    color: 'grey',
    padding: theme.spacing(0, 2),
    height: '100%',
    position: 'absolute',
    pointerEvents: 'none',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  inputRoot: {
    color: 'black',
    width: '100%',
  },
  inputInput: {
    padding: theme.spacing(1, 1, 1, 0),
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
    paddingRight: `calc(1em + ${theme.spacing(4)}px)`,
    transition: theme.transitions.create('width'),
    width: '100%',
  },
}));

const Navbar = () => {
  const classes = useStyles();
  const [{ basket, user }, dispatch] = useStateValue();
  const [searchInput, setSearchInput] = useState("");
  const navigate = useNavigate();

  //manejo de auth
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
    }
  };

  const handleClick = () => {
    const url = `/reload/${searchInput}`;
    navigate(url);
  }

  const reload = () => {
    navigate('/reload');
  }

  return (
    <>
      <CssBaseline />
      <div className={classes.root}>
        <AppBar position='fixed' className={classes.appBar}>
          <Toolbar>
            <Link to='/' className={classes.link}>
              <IconButton>
                <Typography variant="h5" color="textPrimary">
                  E-Commerce
                </Typography>
              </IconButton>
            </Link>
            
            <div className={classes.grow} />

            <div className={classes.search}>
              <div className={classes.searchIcon}>
                <SearchIcon fontSize='mid' />
              </div>
              <div>
                <InputBase
                  classes={{
                    root: classes.inputRoot,
                    input: classes.inputInput,
                  }}
                  inputProps={{ 'aria-label': 'search' }}
                  onChange={(e) => setSearchInput(e.target.value)}
                />
              </div>
            </div>
            <div>
                <Button variant='outlined' onClick={handleClick}>
                  search
                </Button>
            </div>
            <div className={classes.grow} />
              <Typography variant='h6' color='textPrimary' component='p'>
                Hello {user ? user.username : "Guest"}
              </Typography>
              {user? <Menu></Menu> : ""}
              <div className={classes.button}>
                {user? "" : <Link to={!user && "/login"} className={classes.link}>
                  <Button onClick={handleAuth} variant='outlined'>
                    <strong>LogIn</strong>
                  </Button>
                </Link>}
                {!user? "" : <Link to='/checkout' className={classes.link}>
                  <IconButton aria-label='show cart items' color='inherit'>
                    <Badge badgeContent={getItemsQuantity(basket)} color='secondary'>
                      <ShoppingCart fontSize='large' color='primary' />
                    </Badge>
                  </IconButton>
                </Link>}
            </div>
          </Toolbar>
        </AppBar>
      </div>
    </>
  );
};

export default Navbar;