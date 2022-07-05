import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import IconButton from '@material-ui/core/IconButton';
import { Badge, Button, CssBaseline } from '@material-ui/core';
import { ShoppingCart } from '@material-ui/icons';
import { Link } from 'react-router-dom';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    marginBottom: '7rem',
  },
  appBar: {
    backgroundColor: 'whitesmoke',
    boxShadow: 'none',
  },
  grow: {
    flexGrow: 1,
  },
  button: {
    marginLeft: theme.spacing(2),
  },
  image: {
    marginRight: '10px',
    height: '2.5rem',
  },
}));

const Navbar = () => {
  const classes = useStyles();

  return (
    <>
      <CssBaseline />
      <div className={classes.root}>
        <AppBar position='fixed' className={classes.appBar}>
          <Toolbar>
            <Link to='/'>
              <IconButton>
                <Typography variant='h5' color='textPrimary'>
                  E-Commerce
                </Typography>
              </IconButton>
            </Link>
            
            <div className={classes.grow} />
            <Typography variant='h6' color='textPrimary' component='p'>
              Hello guest
            </Typography>
            <div className={classes.button}>
            
                <Button variant='outlined'>
                  <strong>Login</strong>
                </Button>
                
              <Link to='/checkout'>
                <IconButton aria-label='show cart items' color='inherit'>
                  <Badge badgeContent={1} color='secondary'>
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