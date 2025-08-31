import { render } from 'solid-js/web'
import {Route, Router} from "@solidjs/router";
import {lazy} from "solid-js";
import type {JSX} from "solid-js";
import Home, {AppLayout} from "./Home.tsx";


const root = document.getElementById('root')

if (!root) {
    throw new Error("Root div not found")
}



function RouteTree(): JSX.Element{
    return (
        <Router>
            <Route path={"/"} component={AppLayout}>
                <Route path={""} component={Home}/>
            </Route>
            <Route path={"/user"}>
                <Route path={"/login"} component={lazy(()=>import("./pages/Login.tsx"))}></Route>
                <Route path={"/register"} component={lazy(()=>import("./pages/Register.tsx"))}></Route>
            </Route>
        </Router>
    )
}

render(RouteTree, root!)
