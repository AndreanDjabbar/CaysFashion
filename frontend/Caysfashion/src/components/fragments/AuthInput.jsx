import { Input } from "../elements/Input";
import { Label } from "../elements/Typography";

export const AuthInput = (props) => {
    return (
        <div className={props.className}>
            <Label htmlFor={props.id}>{props.children}</Label>
            <Input
            type={props.type}
            id={props.id}
            name={props.name}
            placeholder={props.placeholder}/>
        </div>
    )
}