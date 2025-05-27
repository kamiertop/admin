import './App.css';
import {JSX} from "solid-js";
import {Route, Router} from "@solidjs/router";
import Login from "./pages/Login.tsx";
import Home from "./pages/Home.tsx";
import Register from "./pages/Register.tsx";


export default function App() {
    return (
        <RouteTree/>
    )
}

function RouteTree(): JSX.Element {
    return (
        <Router>
            <Route path={"/user"}>
                <Route path={"/login"} component={Login}/>
                <Route path={"/register"} component={Register}/>
            </Route>
            <Route path={"/"} component={Home}/>
        </Router>
    )
}