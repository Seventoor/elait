import * as Flags from 'country-flag-icons/react/3x2'

export default function FlagIcon ({ code }: { code: string }) {
    const country = code.split('-').pop()?.toUpperCase()
    const Flag = Flags[country as keyof typeof Flags]

    return Flag ? <Flag width={20} /> : <span>🌐</span>
}