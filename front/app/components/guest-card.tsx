import { Form, Link } from "@remix-run/react";
import { BusIcon, PenIcon, TrashIcon, VeggieIcon } from "~/icons";
import { Guest } from "~/lib/models";

interface Props {
  guest: Guest;
}

export function GuestCard({ guest }: Props) {
  return (
    <div className="my-1 flex w-full flex-col rounded-md border-black bg-gradient-to-r from-sky-900 to-indigo-500 px-6 py-2">
      <div className="flex flex-row items-center justify-between">
        <h2 className="text-xl text-white">{guest.name}</h2>
        <div className="flex flex-row gap-4">
          <Link
            className="rounded-full bg-white px-2 py-1"
            to={`/profile/guest?id=${guest.id}`}
          >
            <PenIcon className="size-5 fill-indigo-500 stroke-white stroke-1" />
          </Link>
          <Form action={`/profile/info?id=${guest.id}`} method="delete">
            <button type="submit" className="rounded-full bg-red-800 px-2 py-1">
              <TrashIcon className="size-5 stroke-white stroke-[1px]" />
            </button>
          </Form>
        </div>
      </div>
      {guest.phone && (
        <p className="text-xs italic text-gray-500">{guest.phone}</p>
      )}
      {guest.allergies.length > 0 && (
        <div className="mt-2 flex flex-row rounded-md bg-sky-300 px-4 font-bold text-black">
          <p>Alergias: </p>
          <span className="ml-2 italic">{guest.allergies.join(", ")}</span>
        </div>
      )}
      <Badges veggie={guest.is_vegetarian} transport={guest.needs_transport} />
    </div>
  );
}

function Badges({
  veggie,
  transport,
}: {
  veggie: boolean;
  transport: boolean;
}) {
  const show_section = veggie || transport ? "flex" : "hidden";
  const show_veggie = veggie ? "flex" : "hidden";
  const show_transport = transport ? "flex" : "hidden";

  return (
    <div className={`my-4 ${show_section} flex-row justify-around font-bold`}>
      <div
        className={`${show_veggie} flex-row items-baseline gap-2 rounded-full bg-rose-800 px-10 py-2 duration-1000 animate-in fade-in zoom-in-125`}
      >
        <span className={`text-sm italic text-white`}>Menu</span>
        <VeggieIcon className={`size-4 fill-white stroke-white stroke-2`} />
      </div>
      <div
        className={`${show_transport} flex-row items-baseline gap-2 rounded-full bg-rose-900 px-10 py-2 duration-1000 animate-in fade-in zoom-in-125`}
      >
        <span className={`text-sm italic text-white`}>Usa</span>
        <BusIcon className={`size-4 stroke-white stroke-2`} />
      </div>
    </div>
  );
}
