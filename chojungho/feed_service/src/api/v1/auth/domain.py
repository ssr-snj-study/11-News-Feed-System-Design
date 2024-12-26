from dataclasses import dataclass, asdict


@dataclass
class AuthDomain:
    ...

    @classmethod
    def from_dict(cls, d):
        return cls(**d)

    def to_dict(self):
        return asdict(self)

    def delete_to_dict_none_data(self):
        return {k: v for k, v in self.to_dict().items() if v is not None}
