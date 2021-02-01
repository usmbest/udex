import React from 'react';

const SvgDgd = props => (
  <svg width={props.width || 64} height={props.height || 64} {...props}>
    <path
      fillRule="evenodd"
      d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16-7.163 16-16 16zM5.5 11v10h10V11h-10zm7 3v-3h1v4h-8v-1h7zm-5 3h6v2h-6v-2zm19-4v-2h-10v10h10v-6h-6v2h4v2h-6v-6h8z"
    />
  </svg>
);

export default SvgDgd;
