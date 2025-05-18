import {A} from "@solidjs/router"
import {JSX} from "solid-js";


export default function NotFound(): JSX.Element{
  return (
    <div>
      <h1>404</h1>
      <p>This page could not be found.</p>
      <A href="/home">Back To Home</A>
    </div>
  )
}
