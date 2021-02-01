import React from 'react';

const SvgNas = props => (
  <svg width={props.width || 64} height={props.height || 64} {...props}>
    <defs>
      <linearGradient id="nas_svg__c" x1="50%" x2="50%" y1="0%" y2="100%">
        <stop offset="0%" stopColor="#FFF" stopOpacity={0.5} />
        <stop offset="100%" stopOpacity={0.5} />
      </linearGradient>
      <circle id="nas_svg__b" cx={16} cy={15} r={15} />
      <filter id="nas_svg__a" width="111.7%" height="111.7%" x="-5.8%" y="-4.2%" filterUnits="objectBoundingBox">
        <feOffset dy={0.5} in="SourceAlpha" result="shadowOffsetOuter1" />
        <feGaussianBlur in="shadowOffsetOuter1" result="shadowBlurOuter1" stdDeviation={0.5} />
        <feComposite in="shadowBlurOuter1" in2="SourceAlpha" operator="out" result="shadowBlurOuter1" />
        <feColorMatrix in="shadowBlurOuter1" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.199473505 0" />
      </filter>
      <path
        id="nas_svg__e"
        d="M12.73 17.655l3.146 6.382 1.616-4.773-4.761-1.609m-5.754-2.769l4.764 1.617 1.618-4.77-6.382 3.153m6.031 2.024c.075.04.106.064.141.076 1.536.522 3.07 1.048 4.61 1.557.125.04.303.006.426-.054 1.761-.859 3.517-1.728 5.273-2.595.5-.247.998-.497 1.543-.769-.12-.046-.182-.072-.247-.094-1.487-.505-2.974-1.013-4.466-1.504a.698.698 0 0 0-.477.03c-1.872.91-3.737 1.835-5.603 2.756-.389.192-.776.386-1.2.597m6.499-4.065c-1.13-2.294-2.246-4.554-3.39-6.875l-3.48 10.267 6.87-3.392M5 15.002c1.094-.541 2.126-1.055 3.16-1.565 1.758-.867 3.516-1.735 5.277-2.596a.652.652 0 0 0 .364-.414c.688-2.053 1.386-4.102 2.082-6.152.026-.076.057-.15.105-.275.07.136.122.232.17.329 1.345 2.73 2.688 5.46 4.044 8.184.066.132.24.242.389.293 2.036.702 4.076 1.392 6.115 2.085.077.026.153.054.294.104-.188.095-.327.168-.467.237-2.628 1.297-5.255 2.597-7.887 3.885a.88.88 0 0 0-.493.565c-.679 2.034-1.375 4.062-2.066 6.092-.019.056-.042.11-.087.226-.177-.355-.33-.658-.48-.962-1.224-2.482-2.449-4.964-3.666-7.45a.718.718 0 0 0-.454-.405c-2.042-.683-4.08-1.38-6.119-2.073-.075-.026-.148-.057-.281-.108"
      />
      <filter id="nas_svg__d" width="115.9%" height="115.9%" x="-8%" y="-5.7%" filterUnits="objectBoundingBox">
        <feOffset dy={0.5} in="SourceAlpha" result="shadowOffsetOuter1" />
        <feGaussianBlur in="shadowOffsetOuter1" result="shadowBlurOuter1" stdDeviation={0.5} />
        <feColorMatrix in="shadowBlurOuter1" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.204257246 0" />
      </filter>
    </defs>
    <g fill="none" fillRule="evenodd">
      <use fill="#000" filter="url(#nas_svg__a)" xlinkHref="#nas_svg__b" />
      <use fill="#222" xlinkHref="#nas_svg__b" />
      <use
        fill="url(#nas_svg__c)"
        style={{
          mixBlendMode: 'soft-light',
        }}
        xlinkHref="#nas_svg__b"
      />
      <circle cx={16} cy={15} r={14.5} stroke="#000" strokeOpacity={0.097} />
      <use fill="#000" filter="url(#nas_svg__d)" xlinkHref="#nas_svg__e" />
      <use fill="#FFF" xlinkHref="#nas_svg__e" />
    </g>
  </svg>
);

export default SvgNas;
