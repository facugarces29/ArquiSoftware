import { makeStyles } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import IconButton from "@material-ui/core/IconButton";
import { Badge, Button, CssBaseline } from "@material-ui/core";
import { ShoppingCart } from "@material-ui/icons";
import { useStateValue } from "../StateProvider";
import { Link, useNavigate } from "react-router-dom";
import { actionTypes } from "../reducer";
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
}));




const Navbar = () => {
  const classes = useStyles();
  const [{ basket, user }, dispatch] = useStateValue();
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
      navigate('/');
    }
  };

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

              <Link to='/checkout-page' className={classes.link}>
                <IconButton aria-label='show cart items' color='inherit'>
                  <Badge badgeContent={basket?.length} color='secondary'>
                    <ShoppingCart fontSize='large' color='primary' />
                  </Badge>
                </IconButton>
              </Link>
            </div>
          </Toolbar>
        </AppBar>
      </div>
    </>
  );
};

export default Navbar;