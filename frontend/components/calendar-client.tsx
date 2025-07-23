'use client'

import { Calendar } from "@/components/ui/calendar"
import type { Dispatch, SetStateAction } from "react"

interface CalendarClientProps {
    selected: Date
    onChange: (date: Date) => void // ✅ fix: terima Date langsung
}

export default function CalendarClient({ selected, onChange }: CalendarClientProps) {
    return (
        <Calendar
            mode="single"
            selected={selected}
            onSelect={(date) => {
                if (date instanceof Date) {
                    onChange(date)
                }
            }}
            className="rounded-md border"
        />
    )
}
