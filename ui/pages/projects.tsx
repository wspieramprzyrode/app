import Container from "@material-ui/core/Container";
import Layout from "../components/layout";
import { Typography, Paper, makeStyles } from "@material-ui/core";
const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(2),
    display: "flex",
    overflow: "auto",
    flexDirection: "column",
  },
}));
const ProjectsPage = () => {
  const classes = useStyles();
  return (
    <Layout title="homepage">
      <Paper className={classes.paper}>
        <Typography variant="h4">Nasze projekty</Typography>
        <Typography variant="body1">Inwentaryzacja budek lęgowych ptaków</Typography>
      </Paper>
    </Layout>
  );
};
export default ProjectsPage;
