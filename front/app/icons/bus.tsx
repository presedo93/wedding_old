interface Props {
  className?: string;
}

export const BusIcon = ({ className }: Props) => {
  return (
    <svg
      fill="none"
      viewBox="0 0 24 16"
      transform="matrix(-1, 0, 0, 1, 0, 0)"
      strokeWidth={1.5}
      stroke="currentColor"
      className={className}
    >
      <g>
        <path d="M4.37,17.74H1.5V6.26A1.92,1.92,0,0,1,3.41,4.35H19.92A1.93,1.93,0,0,1,21.82,6l.72,5.24v6.5H19.67"></path>
        <line x1="15.85" y1="17.74" x2="8.2" y2="17.74"></line>
        <circle cx="17.76" cy="17.74" r="1.91"></circle>
        <circle cx="6.28" cy="17.74" r="1.91"></circle>
        <line x1="4.37" y1="12" x2="23.5" y2="12"></line>
        <line x1="17.76" y1="7.22" x2="17.76" y2="12"></line>
        <line x1="12.02" y1="7.22" x2="12.02" y2="12"></line>
        <line x1="6.28" y1="7.22" x2="6.28" y2="12"></line>
      </g>
    </svg>
  );
};
