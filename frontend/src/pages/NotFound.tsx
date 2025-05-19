import {A} from "@solidjs/router";
import type {JSX} from "solid-js";

export default function NotFound(): JSX.Element {
    return (
        <div>
            <h1>404</h1>
            <p>Oops! This page could not be found.</p>
            <A href="/home">
                Back To Home
            </A>
        </div>
    );
}
