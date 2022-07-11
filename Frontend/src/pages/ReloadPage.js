import { makeStyles } from "@material-ui/core/styles";
import { useParams, useNavigate } from "react-router-dom";
import { useEffect } from "react";

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
}));

const Reload = () => {
  const classes = useStyles();
  const param = useParams();
  const navigate = useNavigate();
  const url = `/search/${param.id}`;

  useEffect(() => {
    navigate(url);
  }, []);
  return (
    <>
    </>
  );
};

export default Reload;