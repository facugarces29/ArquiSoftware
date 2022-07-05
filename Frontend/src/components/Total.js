import { Button, makeStyles } from "@material-ui/core";
import accounting from "accounting";
//import { useStateValue } from "../StateProvider";
//import { getBasketTotal } from "../reducer";
//import { Link } from "react-router-dom";
import products from "../product-data";

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
  //const [{ basket }, dispatch] = useStateValue();

  return (
    <div className={classes.root}>
      <h5>Total items : {products.length}</h5>
      <h5>{accounting.formatMoney(50,"$")}</h5>
      <Button
        //component={Link}
        //to='/checkout'
        className={classes.button}
        variant='contained'
        color='secondary'
      >
        Check out
      </Button>
    </div>
  );
};

export default Total;