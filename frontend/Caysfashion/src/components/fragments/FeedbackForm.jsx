import React from "react";
import {Input} from "../elements/Input";
import {Button} from "../elements/Button";
import { Title, Paragraph } from "../elements/Typography";
import { Label } from "../elements/Typography";

export const FeedbackForm = () => {
    return (
        <form>
            <div className="feedback-header">
                <Label>Drop your feedback</Label>
                <Paragraph>Your Feedback Fuels Our Improvement</Paragraph>
            </div>
            <div className="feedback-input">
                <Input 
                id="feedback" 
                name="feedback" placeholder="Feedback"/>
                <Button>Submit</Button>
            </div>
        </form>
    )
};


