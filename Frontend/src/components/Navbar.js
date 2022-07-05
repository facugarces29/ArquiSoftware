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

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    marginBottom: "7rem",
  },
  appBar: {
    backgroundColor: "whitesmoke",
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
}));




const Navbar = () => {
  const classes = useStyles();
  const [{ basket, user }, dispatch] = useStateValue();
  const history = useNavigate();


  //manejo de auth
  const handleAuth = () => {
    if (user) {
      //deslogear
      dispatch({
        type: actionTypes.EMPTY_BASKET,
        basket: [],
      });
      history.push("/");
    }
  };

  return (
    <>
      <CssBaseline />
      <div className={classes.root}>
        <AppBar position='fixed' className={classes.appBar}>
          <Toolbar>
            <Link to='/'>
              <IconButton>
                <Typography variant="h5" color="textPrimary">
                  E-Commerce
                </Typography>
              </IconButton>
            </Link>

            <div className={classes.grow} />
            <Typography variant='h6' color='textPrimary' component='p'>
              Hello {user ? user.email : "Guest"}
            </Typography>
            <div className={classes.button}>
              
              <Link to={!user && "/login"}>
                <Button onClick={handleAuth} variant='outlined'>
                  <strong>{user ? "LogOut" : "LogIn"}</strong>
                </Button>
              </Link>

              <Link to='/checkout-page'>
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