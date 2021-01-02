import React, {PureComponent} from "react";

import altImg from "./../../resource/smantura.jpg"

export default class Image extends PureComponent<any, {errored: boolean, src: any}>
{
    constructor(props)
    {
        super(props);

        this.state = {
            src: props.src ?? (props.fallbackSrc ?? altImg),
            errored: false,
        };
    }

    onError = () =>
    {
        if (!this.state.errored)
        {
            this.setState({
                src: this.props.fallbackSrc ?? altImg,
                errored: true,
            });
        }
    }

    render()
    {
        const { src } = this.state;
        const {
            src: _1,
            fallbackSrc: _2,
            ...props
        } = this.props;

        return (
            <img
                src={src}
                onError={this.onError}
                alt={props.alt}
                {...props}
            />
        );
    }
}
