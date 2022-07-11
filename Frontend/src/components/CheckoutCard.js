import Card from "@material-ui/core/Card";
import CardHeader from "@material-ui/core/CardHeader";
import CardMedia from "@material-ui/core/CardMedia";
import CardActions from "@material-ui/core/CardActions";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import DeleteIcon from "@material-ui/icons/Delete";
import ButtonGroup from '@material-ui/core/ButtonGroup';
import AddIcon from '@material-ui/icons/Add';
import { useStateValue } from "../StateProvider";
import accounting from "accounting";
import { actionTypes } from "../reducer";
import RemoveIcon from '@material-ui/icons/Remove';
import { makeStyles } from "@material-ui/core";

const useStyles = makeStyles((theme) => ({
  root: {
    minWidth: 300,
  },
  action: {
    marginTop: "1rem",
  },
  media: {
    height: 0,
    paddingTop: "56.25%", // 16:9
  },
  cardActions: {
    display: "flex",
    justifyContent: "space-between",
    textAlign: "center",
  },
  cardRating: {
    display: "flex",
  },
}));

const CheckoutCard = ({ product: { id, name, image, price, category, description , stock , quantity } }) => {
  const classes = useStyles();
  const [{ basket }, dispatch] = useStateValue();
  
  const removeItem = () => {
    dispatch({
      type: actionTypes.REMOVE_ITEM,
      id: id,
    });
  };


  const addOne = () => {
    basket.map((item)=>{
      if (item.id === id && item.quantity < stock){
        dispatch({
          type: actionTypes.INCREMENT_ITEM,
          id: id,
        });
      } else if (item.id === id && item.quantity >= stock){
        return
      }
    });
  }

  const removeOne = () => {
    dispatch({
      type: actionTypes.DECREMENT_ITEM,
      id: id,
    });
  }

  return (
    <Card className={classes.root}>
      <CardHeader
        action={
          <Typography
            className={classes.action}
            variant='h5'
            color='textSecondary'
          >
            {accounting.formatMoney(price, "$")}
          </Typography>
        }
        title={name}
        subheader='in Stock'
      />
      <CardMedia className={classes.media} image={image} title={name} />
      <CardActions disableSpacing className={classes.cardActions}>
        <ButtonGroup disableElevation variant="contained"  style={{ color: 'black' }}>
          <IconButton onClick={addOne}>
            <AddIcon style={{ color: 'black' }}></AddIcon>
          </IconButton>
          <IconButton onClick={removeOne}>
            <RemoveIcon style={{ color: 'red' }}></RemoveIcon>
          </IconButton>
        </ButtonGroup>
        Unidades:
        <Typography variant="h6">{quantity} / {stock}</Typography>
        <IconButton onClick={removeItem}>
          <DeleteIcon fontSize='large' />
        </IconButton>
      </CardActions>
    </Card>
  );
};

export default CheckoutCard;