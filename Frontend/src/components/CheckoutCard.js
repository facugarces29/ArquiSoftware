import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
//import clsx from 'clsx';
import Card from '@material-ui/core/Card';
import CardHeader from '@material-ui/core/CardHeader';
import CardMedia from '@material-ui/core/CardMedia';
import CardActions from '@material-ui/core/CardActions';
//import Collapse from '@material-ui/core/Collapse';
import Typography from '@material-ui/core/Typography';
//import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import accounting from 'accounting'
//import { AddShoppingCart } from '@material-ui/icons';
import DeleteIcon from '@material-ui/icons/Delete';
import { IconButton } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
    root: {
      maxWidth: 345,
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
  }));
  

export default function CheckoutCard({ productName , isInStock , productImage , productPrice }) {
  const classes = useStyles();
  //const [expanded, setExpanded] = React.useState(false);

  /*const handleExpandClick = () => {
    setExpanded(!expanded);
  };*/

  return (
    <Card className={classes.root}>
      <CardHeader
       action={
        <Typography
          className={classes.action}
          variant='h5'
          color='textSecondary'
        >
          {accounting.formatMoney(productPrice)}
        </Typography>
      }
        title={productName}
        subheader={ isInStock ? 'in stock' : 'no stock' }
      />
      <CardMedia 
        className={classes.media} 
        image={productImage} 
        title={productName} />
      <CardActions disableSpacing className={classes.cardActions}>
        <IconButton >
          <DeleteIcon fontSize='large' />
        </IconButton>
      </CardActions>
    </Card>
  );
}