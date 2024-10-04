// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import classNames from 'classnames';
import React from 'react';
import type {MessageDescriptor} from 'react-intl';
import {useIntl} from 'react-intl';

type ButtonProps = {
    onClick: React.MouseEventHandler<HTMLButtonElement>;
    label: string | MessageDescriptor;
    emphasis: 'primary' | 'secondary' | 'tertiary' | 'quaternary' | 'link';
    size?: 'xSmall' | 'small' | 'medium' | 'large';
    pull?: 'left' | 'right';
    destructive?: boolean;
    leadingIcon?: string;
    trailingIcon?: string;
    autoFocus?: boolean;
    testId?: string;
    disabled?: boolean;
}

const emphasisClasses = {
    primary: 'btn-primary',
    secondary: 'btn-secondary',
    tertiary: 'btn-tertiary',
    quaternary: 'btn-quaternary',
    link: 'btn-link',
};

const sizeClasses = {
    xSmall: 'btn-xs',
    small: 'btn-sm',
    medium: '',
    large: 'btn-lg',
};

const Button = ({
    onClick,
    label,
    emphasis = 'primary',
    size = 'medium',
    pull,
    destructive,
    leadingIcon,
    trailingIcon,
    autoFocus,
    testId,
    disabled,
}: ButtonProps) => {
    const intl = useIntl();
    const leading = leadingIcon ? (<i className={classNames('icon', leadingIcon)}/>) : null;
    const trailing = trailingIcon ? (<i className={classNames('icon', trailingIcon)}/>) : null;

    const emphasisClass = emphasisClasses[emphasis];
    const sizeClass = sizeClasses[size];

    const labelText = typeof label === 'string' ? label : intl.formatMessage(label);
    return (
        <button
            type='button'
            onClick={onClick}
            className={classNames('btn', emphasisClass, sizeClass, {
                'pull-left': pull === 'left',
                'pull-right': pull === 'right',
                'btn-danger': destructive,
            })}
            autoFocus={autoFocus}
            data-testid={testId}
            disabled={disabled}
        >
            {leading}
            {labelText}
            {trailing}
        </button>
    );
};

export default Button;
