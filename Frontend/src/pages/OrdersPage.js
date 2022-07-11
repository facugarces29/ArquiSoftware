import { makeStyles } from "@material-ui/core/styles";
import { Grid, CssBaseline } from "@material-ui/core";
import { useState , useEffect} from "react";
import Typography from "@material-ui/core/Typography";
import { useStateValue } from "../StateProvider";
import OrderCard from "../components/OrderCard"

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
  const [address, setAddress] = useState();
  const [{ user }, dispatch] = useStateValue();
  let idUser = 0;

  user? idUser = user.userID : idUser = 0;

  console.log(idUser);
  
  const urlOrders = "http://localhost:8080/order/" + idUser;
  const urlAddress = "http://localhost:8080/address/" + idUser;
  
  const fetchOrders = async () => {
    const response = await fetch(urlOrders);
    const responseJSON = await response.json();
    setOrders(responseJSON);
  }

  const fetchAddress = async () => {
    const response = await fetch(urlAddress);
    const responseJSON = await response.json();
    setAddress(responseJSON);
  }

  useEffect(() => {
    fetchOrders();
    fetchAddress();
  }, []);


  return (
    <>
      <CssBaseline/>
      <div className={classes.root}>
        <Grid container spacing={3} className={classes.root}>
          { !orders ? <Typography  variant='h3' color='textPrimary' className={classes.text}> todavia no has realizado ninguna orden.</Typography> : 
            orders.map((order) => {
            console.log("order: ", order);
            console.log("products: ", order.order_products)
            return  <Grid item xs={12} sm={6} md={4} lg={4} className={classes.card}>
                      <OrderCard
                        date={order.date}
                        amount={order.amount}
                        image={order.image}
                        products = {order.order_products}
                        address = {address}
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