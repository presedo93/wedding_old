import { motion } from "framer-motion";
import { useState } from "react";
import { Checkbox } from "~/components/ui/checkbox";

interface Props {
  children: React.ReactNode;
}

export function TodoItem({ children }: Props) {
  const [checked, setChecked] = useState(false);

  return (
    <li className="flex flex-row items-center gap-2 text-lg font-medium text-white">
      <Checkbox
        className="border-sky-950 data-[state=checked]:bg-sky-950"
        checked={checked}
        onClick={() => setChecked(!checked)}
      />
      <div className="relative">
        <span className="font-sand text-lg">{children}</span>
        <motion.div
          className="absolute left-0 top-1/2 h-0.5 rounded-xl bg-white"
          initial={{ width: checked ? 0 : 1 }}
          animate={{ width: checked ? "100%" : "0%" }}
          transition={{ duration: 0.75 }}
        />
      </div>
    </li>
  );
}
