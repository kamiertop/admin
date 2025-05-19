import {Route, Router} from "@solidjs/router"
import NotFound from "./pages/NotFound.tsx";


export default function App() {
    return (
        <Router>
            <Route path={"/"}/>
            <Route path={"*404"} component={NotFound}/>
        </Router>
    )
}