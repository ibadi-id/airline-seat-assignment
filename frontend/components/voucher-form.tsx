'use client'

import { useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"

import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert"
import { format } from "date-fns"
import dynamic from "next/dynamic"
import React from "react"

const CalendarClient = dynamic(() => import("@/components/calendar-client"), {
    ssr: false,
})

const formSchema = z.object({
    name: z.string().min(1, "Please enter the crew member's full name."),
    crewId: z.string().min(1, "Crew ID is required to identify the crew."),
    flightNumber: z.string().min(1, "Flight number cannot be empty."),
    flightDate: z.date(),
    aircraftType: z.enum(["ATR", "Airbus 320", "Boeing 737 Max"]),
})

type FormData = z.infer<typeof formSchema>

export function VoucherForm() {
    const [assignedSeats, setAssignedSeats] = useState<string[] | null>(null)
    const [error, setError] = useState<string | null>(null)

    const {
        register,
        handleSubmit,
        setValue,
        watch,
        formState: { errors },
    } = useForm<FormData>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            flightDate: new Date(),
        },
    })

    const onSubmit = async (data: FormData) => {
        setError(null)
        setAssignedSeats(null)

        const dateString = format(data.flightDate, "yyyy-MM-dd")

        try {
            const checkRes = await fetch("http://localhost:8080/api/check", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    flight_number: data.flightNumber,
                    date: dateString,
                }),
            })

            const checkJson = await checkRes.json()
            if (checkJson.exists) {
                setError("🚫 Vouchers already generated for this flight/date.")
                return
            }

            const generateRes = await fetch("http://localhost:8080/api/generate", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    name: data.name,
                    id: data.crewId,
                    flight_number: data.flightNumber,
                    date: dateString,
                    aircraft: data.aircraftType,
                }),
            })

            const generateJson = await generateRes.json()

            if (generateJson.success) {
                setAssignedSeats(generateJson.seats)
            } else {
                setError("❌ Failed to generate seats.")
            }
        } catch (err) {
            setError("⚠️ Network or server error. Please try again.")
        }
    }

    return (
        <div className="flex flex-col md:flex-row gap-8 mt-8">
            {/* Form Column */}
            <div className="md:w-1/2 w-full">
                <Card>
                    <CardHeader className="text-center">
                        <CardTitle className="text-xl">🎫 Airline Voucher Seat Assignment</CardTitle>
                        <CardDescription>
                            <p className="text-muted-foreground">
                                Enter your voucher details to automatically assign your seat
                            </p>
                        </CardDescription>
                    </CardHeader>
                    <CardContent>
                        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
                            {/* Name */}
                            <div className="grid gap-2">
                                <Label htmlFor="crewname">Crew Name</Label>
                                <Input
                                    {...register("name")}
                                    className={errors.name ? "border-red-500" : ""}
                                />
                                {errors.name && <p className="text-red-500 text-sm">{errors.name.message}</p>}
                            </div>

                            {/* Crew ID */}
                            <div className="grid gap-2">
                                <Label htmlFor="crewId">Crew ID</Label>
                                <Input
                                    {...register("crewId")}
                                    className={errors.crewId ? "border-red-500" : ""}
                                />
                                {errors.crewId && <p className="text-red-500 text-sm">{errors.crewId.message}</p>}
                            </div>

                            {/* Flight Number */}
                            <div className="grid gap-2">
                                <Label>Flight Number</Label>
                                <Input
                                    {...register("flightNumber")}
                                    className={errors.flightNumber ? "border-red-500" : ""}
                                />
                                {errors.flightNumber && <p className="text-red-500 text-sm">{errors.flightNumber.message}</p>}
                            </div>

                            {/* Flight Date */}
                            <div className="grid gap-2">
                                <Label>Flight Date</Label>
                                <CalendarClient
                                    selected={watch("flightDate")}
                                    onChange={(date) => setValue("flightDate", date)}
                                />
                            </div>

                            {/* Aircraft Type */}
                            <div className="grid gap-2">
                                <Label>Aircraft Type</Label>
                                <Select onValueChange={(v) => setValue("aircraftType", v as any)}>
                                    <SelectTrigger>
                                        <SelectValue placeholder="Select Aircraft Type" />
                                    </SelectTrigger>
                                    <SelectContent>
                                        <SelectItem value="ATR">ATR</SelectItem>
                                        <SelectItem value="Airbus 320">Airbus 320</SelectItem>
                                        <SelectItem value="Boeing 737 Max">Boeing 737 Max</SelectItem>
                                    </SelectContent>
                                </Select>
                                {errors.aircraftType && <p className="text-red-500 text-sm">{errors.aircraftType.message}</p>}
                            </div>

                            <Button type="submit" className="w-full">Generate Vouchers</Button>
                        </form>
                    </CardContent>
                </Card>
            </div>

            {/* Result Column */}
            <div className="md:w-1/2 w-full h-full">
                {/* <div className="flex-1 border rounded-md p-4 bg-muted/50 space-y-2">
                    <h3 className="text-lg font-semibold mb-2">🛫 Your Assigned Seats</h3>
                    {!assignedSeats && (<p className="text-muted-foreground">Please submit the form to see the seat assignment result.</p>)} */}


                <Card >
                    <CardHeader>
                        <CardTitle className="text-lg font-semibold">🛫 Your Assigned Seats</CardTitle>
                        <CardDescription> {!assignedSeats && (<p className="text-muted-foreground">Please submit the form to see the seat assignment result.</p>)} </CardDescription>
                    </CardHeader>
                    {assignedSeats && (
                        <CardContent>
                            <div className="flex flex-wrap gap-4 justify-center">
                                {assignedSeats.map((seat, i) => (
                                    <div
                                        key={i}
                                        className="flex flex-col items-center justify-center w-24 h-24 bg-blue-500 text-white rounded-xl shadow-md"
                                    >
                                        <span className="text-sm font-semibold text-white/80">Seat {i + 1}</span>
                                        {/* <span className="text-sm font-bold text-white/90">{i + 1}</span> */}
                                        <span className="mt-1 text-lg font-extrabold">{seat}</span>
                                    </div>
                                ))}
                            </div>
                            <p className="mt-4 text-center text-gray-600 text-sm">
                                Please proceed to boarding with your assigned seats.
                            </p>
                        </CardContent>
                    )}
                    {error && (
                        <CardContent>
                            <Alert variant="destructive" className="mt-4">
                                <AlertTitle>Error</AlertTitle>
                                <AlertDescription>{error}</AlertDescription>
                            </Alert>
                        </CardContent>
                    )}
                </Card>



                {/* </div> */}

            </div>
        </div>
    )
}
