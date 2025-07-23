import Image from "next/image";
import { VoucherForm } from "@/components/voucher-form"

export default function Home() {
  return (
    <div className="bg-muted min-h-screen flex items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-5xl bg-white rounded-md shadow-md p-6 flex flex-col md:flex-row gap-6">
        <div className="flex-1">
          <a href="#" className="flex items-center gap-2 font-medium mb-4">
            <div className="relative size-8">
              <Image
                src="/logo.jpg"
                alt="Book Cabin Logo"
                fill
                sizes="24px"
                className="object-contain rounded-md"
              />
            </div>
            <span className="text-lg">Book Cabin</span>
          </a>
          <VoucherForm />
        </div>
      </div>
    </div>
  );
}
