import Link from "next/link";
import { Button } from "@/components/ui/button";
import UserItem from "@/components/userItem";

export default async function Home() {

  return (
    <main>
      <div className="flex flex-col space-y-3 text-center p-10 items-center">
        <UserItem />
        <Link href="/apis">
          <Button className="bg-zinc-300 dark:bg-zinc-900 text-black dark:text-white">Go to APIs</Button>
        </Link>
      </div>
    </main>
  );
}
