import { makeStyles } from "@material-ui/core/styles";
import { Grid, CssBaseline } from "@material-ui/core";
import { useState , useEffect} from "react";
import Avatar from "@material-ui/core/Avatar";
import Button from "@material-ui/core/Button";
import TextField from "@material-ui/core/TextField";
import Link from "@material-ui/core/Link";
import Box from "@material-ui/core/Box";
import LockOutlinedIcon from "@material-ui/icons/LockOutlined";
import Typography from "@material-ui/core/Typography";
import Container from "@material-ui/core/Container";
import { actionTypes } from "../reducer";
import { useStateValue } from "../StateProvider";
import { useNavigate } from "react-router-dom";
import OrderCard from "../components/OrderCard"
import { ContactSupportOutlined } from "@material-ui/icons";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
  card: {
    display:'flex',
    alignItems:'center',
    justifyContent:'center',
  },
  text:{
    display:'flex',
    alignItems:'center',
    justifyContent:'center',
  },
}));

const Orders = () => {
  const classes = useStyles();
  const [orders, setOrders] = useState();
  const [{ user }, dispatch] = useStateValue();
  let idUser = 0;

  user? idUser = user.userID : idUser = 0;

  console.log(idUser);
  
  const urlOrders = "http://localhost:8080/order/" + idUser;
  
  const fetchOrders = async () => {
    const response = await fetch(urlOrders);
    const responseJSON = await response.json();
    setOrders(responseJSON);
  }

  useEffect(() => {
    fetchOrders();
  }, []);


  return (
    <>
      <CssBaseline/>
      <div className={classes.root}>
        <Grid container spacing={3} className={classes.root}>
          { !orders ? <Typography  variant='h3' color='textPrimary' className={classes.text}> todavia no has realizado ninguna orden.</Typography> : 
            orders.map((order) => {
            return  <Grid item xs={12} sm={6} md={4} lg={4} className={classes.card}>
                      <OrderCard
                        date={order.date}
                        amount={order.amount}
                        image={order.image}
                      ></OrderCard>
                      </Grid>
            })

          }
          </Grid>
      </div>
    </>
  );
};

export default Orders;