import { useState } from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import CssBaseline from "@material-ui/core/CssBaseline";
import TextField from "@material-ui/core/TextField";
import Link from "@material-ui/core/Link";
import Box from "@material-ui/core/Box";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Typography from "@material-ui/core/Typography";
import { makeStyles } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";
import { actionTypes } from "../reducer";
import { useStateValue } from "../StateProvider";
import { useNavigate } from "react-router-dom"

function Copyright() {
  return (
    <Typography variant='body2' color='textSecondary' align='center'>
      {"Copyright © "}
      <Link color='inherit' href='localhost:3000/'>
        E-Commerce
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
}

const useStyles = makeStyles((theme) => ({
  paper: {
    marginTop: theme.spacing(8),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.secondary.main,
  },
  form: {
    width: "100%", // Fix IE 11 issue.
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
  },
  error:{
    color: "red",
  },
}));

let couldLogin = false;
async function login(username, password) {
    return await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({"username":username, "password":password})
    })
      .then(response => {
        if (response.status === 400 || response.status === 401)
        {
          couldLogin = false;
          return;
        }
        couldLogin = true;
        return;
      })
   }


export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(false);
  const navigate = useNavigate();
  const [{ user }, dispatch] = useStateValue();
  const classes = useStyles();

  const signin = (e) => {
    console.log("entro a signin")
    e.preventDefault();
    login(username,password).then(data => {
      if (couldLogin === true){
        console.log("entro al if")
        logUser();
        navigate('/');
      } else{
        setError(true);
      }
    });
  };

  const logUser = () => {
    dispatch({
      type: actionTypes.SET_USER,
      user: {
        username,
        password,
      },
    });
  };

  return (
    <Container component='main' maxWidth='xs'>
      <CssBaseline />
      <div className={classes.paper}>
        <Avatar className={classes.avatar}>
          <LockOutlinedIcon />
        </Avatar>
        <Typography component='h1' variant='h5'>
          Login
        </Typography>
        <form className={classes.form} noValidate>
          <TextField
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            variant='outlined'
            margin='normal'
            required
            fullWidth
            id='username'
            label='Username'
            name='username'
            autoFocus
          />
          <TextField
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            variant='outlined'
            margin='normal'
            required
            fullWidth
            name='password'
            label='Password'
            type='password'
            id='password'
            autoComplete='current-password'
          />
          <div className={classes.error}>
            <Typography variant="h5">
              {error? "Usuario o contraseña inválidos." : "" }
            </Typography>
          </div>
          <Button
            onClick={signin}
            type='submit'
            fullWidth
            variant='contained'
            color='primary'
            className={classes.submit}
          >
            Login
          </Button>
        </form>
      </div>
      <Box mt={8}>
        <Copyright />
      </Box>
    </Container>
  );
}