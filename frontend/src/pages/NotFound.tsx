import {JSX} from "solid-js";
import {A} from "@solidjs/router";


export default function NotFound(): JSX.Element {
    return (
        <div class="font-extrabold">
            <h1>404NotFound</h1>
            <A href="/">
                Go Home
            </A>
        </div>
    )
}