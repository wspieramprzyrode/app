import React from "react";
import Head from "next/head";
import { ThemeProvider, makeStyles } from "@material-ui/core/styles";
import CssBaseline from "@material-ui/core/CssBaseline";
import theme from "../lib/theme";
import { Header as HeaderComponent } from "./header";
import { Footer as FooterComponent } from "./footer";
import { Container } from "@material-ui/core";
const useStyles = makeStyles((theme) => ({
  root: {
    display: "flex",
  },
  appBarSpacer: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    height: "100vh",
    overflow: "auto",
  },
  container: {
    paddingTop: theme.spacing(4),
    paddingBottom: theme.spacing(4),
  },
  fotter: {
    textAlign: "center",
    position: "absolute",
    bottom: 0,
    width: "100% !important",
    height: "100px !important",
  },
}));
export default function Layout({
  children,
  title,
}: {
  children: React.ReactNode;
  title: string;
}) {
  const classes = useStyles();
  return (
    <React.Fragment>
      <Head>
        <title>WspieramPrzyrode - {title}</title>
      </Head>
      <ThemeProvider theme={theme}>
        <div className={classes.root}>
          <CssBaseline />
          <header>
            <HeaderComponent />
          </header>
          <main className={classes.content}>
            <div className={classes.appBarSpacer} />
            <Container maxWidth="lg" className={classes.container}>
            {children}
            </Container>
          </main>
          <footer className={classes.fotter}>
            <FooterComponent />
          </footer>
        </div>
      </ThemeProvider>
    </React.Fragment>
  );
}
