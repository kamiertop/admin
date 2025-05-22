import {A} from "@solidjs/router";
import {JSX} from "solid-js";


export default function NotFound(): JSX.Element {
    return (
        <div>
            <h1>404NotFound</h1>
            <A href="/">
                Go Home
            </A>
        </div>
    )
}