import { Button, makeStyles } from "@material-ui/core";
import accounting from "accounting";
import { useStateValue } from "../StateProvider";
import { getBasketTotal, getItemsQuantity } from "../reducer";
import { Link } from "react-router-dom";

const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    height: "20vh",
  },
  button: {
    maxWidth: "200px",
    marginTop: "2rem",
  },
}));

const Total = () => {
  const classes = useStyles();
  const [{ basket }, dispatch] = useStateValue();

  let isBasket = false;

  if (getItemsQuantity(basket) > 0) {
    isBasket = true;
  }

  return (
    <div className={classes.root}>
      <h5>Total items : {getItemsQuantity(basket)}</h5>
      <h5>{accounting.formatMoney(getBasketTotal(basket), "$")}</h5>

      {isBasket? 
      <Button
        component={Link}
        to='/place-order'
        className={classes.button}
        variant='contained'
        color='secondary'
      >
        Check out
      </Button> : console.log("no hay items")}
    </div>
  );
};

export default Total;